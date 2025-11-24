export type AlbumFormItemProps = {
  id?:number
  name:string
  description?:string
  sort:number
}

export type AlbumPhotoFormItemProps = {
  id?:number
  name:string
  image_url:string
  description?:string
  view_count?:number
  album_id:number
}

export type AlbumFormProps = {
  formInline: AlbumFormItemProps
}

export type AlbumPhotoFormProps ={
  formInline: AlbumPhotoFormItemProps
}
