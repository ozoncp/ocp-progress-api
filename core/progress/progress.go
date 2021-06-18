package progress

type Progress struct {
	Id             uint64 `db:"id"`
	ClassroomId    uint   `db:"classroom_id"`
	PresentationId uint   `db:"presentation_id"`
	SlideId        uint   `db:"slide_id"`
	UserId         uint   `db:"user_id"`
}
