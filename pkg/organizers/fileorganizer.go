package organizers

type FileOrganizer interface {
	GetNewFileFolderMapping(workdingDir string) (map[string]string, error)
}
