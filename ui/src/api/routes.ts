import axios from "axios"
import { BASE_URL } from "./utils"

export const getAsyncRoutes=()=>{
  return axios.get(`${BASE_URL}/v1/routes`)
}
