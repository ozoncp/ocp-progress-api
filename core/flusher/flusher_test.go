package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	opentracing "github.com/opentracing/opentracing-go"
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
		toFlush       []progress.Progress
		failedToFlush []progress.Progress
		ctx           context.Context
		chSize        int
	)

	BeforeEach(func() {
		ctx = context.TODO()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

	})

	JustBeforeEach(func() {

		tracer := opentracing.GlobalTracer()
		span := tracer.StartSpan("JustBeforeEach")
		defer span.Finish()

		f = flusher.New(mockRepo, chSize)
		failedToFlush = f.Flush(ctx, span, toFlush)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Zero input data", func() {
		BeforeEach(func() {
			chSize = 2
			toFlush = []progress.Progress{{}}

			mockRepo.EXPECT().AddProgress(ctx, gomock.Any()).Return(nil).MinTimes(1)
		})
		It("Rez", func() {
			//Expect(err).Should(BeNil())
			Expect(failedToFlush).Should(BeEmpty())
		})
	})

	Context("Errors in AddProgress", func() {
		BeforeEach(func() {
			chSize = 3

			toFlush = []progress.Progress{
				{Id: 1, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 2, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 3, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 4, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
				{Id: 5, ClassroomId: 1, PresentationId: 1, SlideId: 1, UserId: 1},
			}

			failedToFlush = toFlush
			toFlush = []progress.Progress{{}}

			mockRepo.EXPECT().AddProgress(ctx, gomock.Any()).Return(errors.New("add prize error")).MinTimes(1)
		})
		It("Errors", func() {
			//Expect(err).Should(BeNil())
			Expect(failedToFlush).Should(BeEquivalentTo(toFlush))
		})
	})
})
