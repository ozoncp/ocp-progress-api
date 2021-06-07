package alarmer_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-progress-api/core/alarmer"
)

var _ = Describe("Alarmer", func() {
	var (
		alrm  alarmer.Alarmer
		timer *time.Timer
	)

	Context("Alarm test", func() {
		BeforeEach(func() {
			alrm = alarmer.New(20 * time.Millisecond)
			timer = time.NewTimer(30 * time.Millisecond)
		})

		It("closing alarm", func() {
			alrm.Init()
			timer.Reset(0)

			go func() {
				defer alrm.Close()
				<-timer.C
			}()

			Eventually(alrm.Alarm()).Should(BeClosed())
		})

	})

})
