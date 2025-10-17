import axios from "axios"

export const getAsyncRoutes=()=>{
  return axios.get('/routes')
}
