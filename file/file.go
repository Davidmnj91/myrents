package file

/*File struct representation of an uploaded file in the filesystem*/
type File struct {
	Name string `bson:"name,omitempty" json:"name"`
	FilePath string `bson:"filepath,omitempty" json:"filepath"`
}