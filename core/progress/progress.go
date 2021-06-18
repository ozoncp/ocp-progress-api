package progress

import desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"

type Progress struct {
	Id             uint64 `db:"id"`
	ClassroomId    uint   `db:"classroom_id"`
	PresentationId uint   `db:"presentation_id"`
	SlideId        uint   `db:"slide_id"`
	UserId         uint   `db:"user_id"`
}

func (pr *Progress) ToProtoClassroom() *desc.Progress {

	return &desc.Progress{

		Id:             pr.Id,
		ClassroomId:    uint64(pr.ClassroomId),
		PresentationId: uint64(pr.PresentationId),
		SlideId:        uint64(pr.SlideId),
		UserId:         uint64(pr.UserId),
	}
}
