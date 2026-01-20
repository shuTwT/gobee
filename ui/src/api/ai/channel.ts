import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";
import axios from "axios";

export interface Model {
    id: string,
    object: 'model',
    created: number,
    owned_by: string
}
interface ModelListResult {
    object: 'list',
    data: Model[]
}

export const getModelList = (channelUrl: string,token:string) => {
  return axios.get<ModelListResult>(`${channelUrl}/v1/models`,{
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
}