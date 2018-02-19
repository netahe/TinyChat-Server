package messanger

type DBObject interface {
  Create()
  Read()
  Update()
  Delete()
}
