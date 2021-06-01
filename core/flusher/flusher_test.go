package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-progress-api/core/flusher"
	"github.com/ozoncp/ocp-progress-api/core/mocks"
	"github.com/ozoncp/ocp-progress-api/core/progress"
)

var _ = Describe("Flusher", func() {
	var (
		//err           error
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		f             flusher.Flusher
		toFlush       []progress.Pogress
		failedToFlush []progress.Pogress
		chSize        int
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		/*
			toFlush = []progress.Pogress{
				{Id: 1, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 2, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 3, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 4, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 5, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
			}
		*/
	})

	JustBeforeEach(func() {
		f = flusher.New(mockRepo, chSize)
		failedToFlush = f.Flush(toFlush)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Flushed successfully", func() {
		BeforeEach(func() {
			chSize = 2
			toFlush = []progress.Pogress{{}}

			mockRepo.EXPECT().AddProgress(gomock.Any()).Return(nil).MinTimes(1)
		})
		It("", func() {
			//Expect(err).Should(BeNil())
			Expect(failedToFlush).Should(BeEmpty())
		})
	})
})
