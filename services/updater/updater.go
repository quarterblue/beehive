package updater

type Updater struct {
	jobCount uint64
}

func NewUpdater() *Updater {
	return &Updater{
		jobCount: 0,
	}
}

func (u *Updater) JobUpdate() {}
