package progress

type Pogress struct {
	Id             uint64
	ClassroomId    uint
	PresentationId uint
	SlideId        uint
	UserId         uint
}

func New() *Pogress {
	return &Pogress{0, 0, 0, 0, 0}
}

func MustNew() Pogress {
	return Pogress{0, 0, 0, 0, 0}
}

func (p *Pogress) Init(id uint64, classroomId uint, presentationId uint, slideId uint, userId uint) {
	p.Id = id
	p.ClassroomId = classroomId
	p.PresentationId = presentationId
	p.SlideId = slideId
	p.UserId = userId
}
