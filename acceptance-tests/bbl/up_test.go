package acceptance_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	acceptance "github.com/cloudfoundry/bosh-bootloader/acceptance-tests"
	"github.com/cloudfoundry/bosh-bootloader/acceptance-tests/actors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("up", func() {
	var (
		bbl             actors.BBL
		boshcli         actors.BOSHCLI
		directorAddress string
		caCertPath      string
		sshSession      *gexec.Session
		stateDir        string
		iaas            string

		boshDirectorChecker actors.BOSHDirectorChecker
	)

	BeforeEach(func() {
		acceptance.SkipUnless("bbl-up")

		configuration, err := acceptance.LoadConfig()
		Expect(err).NotTo(HaveOccurred())

		iaas = configuration.IAAS

		stateDir = configuration.StateFileDir

		bbl = actors.NewBBL(stateDir, pathToBBL, configuration, "up-env")
		boshcli = actors.NewBOSHCLI()
		boshDirectorChecker = actors.NewBOSHDirectorChecker(configuration)
	})

	AfterEach(func() {
		if sshSession != nil {
			sshSession.Interrupt()
			Eventually(sshSession, "5s").Should(gexec.Exit())
		}
		session := bbl.Down()
		Eventually(session, 10*time.Minute).Should(gexec.Exit())
	})

	It("bbl's up a new bosh director and jumpbox", func() {
		session := bbl.Up("--name", bbl.PredefinedEnvID())
		Eventually(session, 40*time.Minute).Should(gexec.Exit(0))

		By("verifying that artifacts are created in state dir", func() {
			checkExists := func(dir string, filenames []string) {
				for _, f := range filenames {
					_, err := os.Stat(filepath.Join(dir, f))
					Expect(err).NotTo(HaveOccurred())
				}
			}

			checkExists(stateDir, []string{"bbl-state.json"})
			checkExists(stateDir, []string{"create-jumpbox.sh"})
			checkExists(stateDir, []string{"create-director.sh"})
			checkExists(stateDir, []string{"delete-jumpbox.sh"})
			checkExists(stateDir, []string{"delete-director.sh"})
			checkExists(filepath.Join(stateDir, ".bbl", "cloudconfig"), []string{
				"cloud-config.yml",
				"ops.yml",
			})
			checkExists(filepath.Join(stateDir, "bosh-deployment"), []string{
				"bosh.yml",
				filepath.Join(iaas, "cpi.yml"),
				"credhub.yml",
				"jumpbox-user.yml",
				"uaa.yml",
			})
			checkExists(filepath.Join(stateDir, "jumpbox-deployment"), []string{
				filepath.Join(iaas, "cpi.yml"),
				"jumpbox.yml",
			})
			checkExists(filepath.Join(stateDir, "terraform"), []string{
				"template.tf",
			})
			checkExists(filepath.Join(stateDir, "vars"), []string{
				"bosh-state.json",
				"director-variables.yml",
				"jumpbox-state.json",
				"jumpbox-variables.yml",
				"terraform.tfstate",
				"user-ops-file.yml",
			})
		})

		By("creating an ssh tunnel to the director in print-env", func() {
			sshSession = bbl.StartSSHTunnel()
		})

		By("checking if the bosh director exists", func() {
			directorAddress = bbl.DirectorAddress()
			caCertPath = bbl.SaveDirectorCA()

			directorExists := func() bool {
				exists, err := boshcli.DirectorExists(directorAddress, caCertPath)
				if err != nil {
					fmt.Println(string(err.(*exec.ExitError).Stderr))
				}
				return exists
			}
			Eventually(directorExists, "1m", "10s").Should(BeTrue())
		})

		By("checking that the cloud config exists", func() {
			directorUsername := bbl.DirectorUsername()
			directorPassword := bbl.DirectorPassword()

			cloudConfig, err := boshcli.CloudConfig(directorAddress, caCertPath, directorUsername, directorPassword)
			Expect(err).NotTo(HaveOccurred())
			Expect(cloudConfig).NotTo(BeEmpty())
		})

		By("checking if bbl print-env prints the bosh environment variables", func() {
			stdout := bbl.PrintEnv()

			Expect(stdout).To(ContainSubstring("export BOSH_ENVIRONMENT="))
			Expect(stdout).To(ContainSubstring("export BOSH_CLIENT="))
			Expect(stdout).To(ContainSubstring("export BOSH_CLIENT_SECRET="))
			Expect(stdout).To(ContainSubstring("export BOSH_CA_CERT="))
		})

		By("rotating the jumpbox's ssh key", func() {
			sshKey := bbl.SSHKey()
			Expect(sshKey).NotTo(BeEmpty())

			session = bbl.Rotate()
			Eventually(session, 40*time.Minute).Should(gexec.Exit(0))

			rotatedKey := bbl.SSHKey()
			Expect(rotatedKey).NotTo(BeEmpty())
			Expect(rotatedKey).NotTo(Equal(sshKey))
		})

		By("checking bbl up is idempotent", func() {
			session := bbl.Up()
			Eventually(session, 40*time.Minute).Should(gexec.Exit(0))
		})

		By("destroying the director and the jumpbox", func() {
			session := bbl.Down()
			Eventually(session, 10*time.Minute).Should(gexec.Exit(0))
		})

		By("verifying that artifacts are removed from state dir", func() {
			f, err := os.Open(stateDir)
			Expect(err).NotTo(HaveOccurred())

			filenames, err := f.Readdirnames(0)
			Expect(err).NotTo(HaveOccurred())
			Expect(filenames).To(BeEmpty())
		})
	})
})
