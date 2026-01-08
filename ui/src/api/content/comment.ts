import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

// 评论状态
export enum CommentStatus {
  Pending = 1,  // 未审核
  Approved = 2  // 已审核
}

// 评论列表项类型
export interface Comment {
  id: number;
  post_id: number;
  page_id: number;
  content: string;
  user_id: number;
  status: CommentStatus;
  user_agent?: string;
  ip_address: string;
  ip_location?: string;
  pinned: boolean;
  created_at: string;
  updated_at: string;
  // 关联数据
  user?: {
    id: number;
    name: string;
    email: string;
  };
  article?: {
    id: number;
    title: string;
  };
}

// 获取评论列表
export const getCommentList = (params: {
  page?: number;
  pageSize?: number;
  status?: CommentStatus;
}) => {
  return http.request<ApiResponse<Comment>>('get',`${BASE_URL}/v1/comment`, { params });
};

// 获取评论分页列表
export const getCommentPage = (params: {
  page?: number;
  pageSize?: number;
  status?: CommentStatus;
}) => {
  return http.request<TableResponse<Comment>>('get',`${BASE_URL}/v1/comment/page`, { params });
};

// 获取评论详情
export const getCommentDetail = (id: number) => {
  return http.request<ApiResponse<Comment>>('get',`${BASE_URL}/v1/comment/query/${id}`);
};

// 审核评论
export const approveComment = (id: number) => {
  return http.request<ApiResponse<Comment>>('put',`${BASE_URL}/v1/comment/approve/${id}`);
};

// 拒绝评论
export const rejectComment = (id: number) => {
  return http.request<ApiResponse<Comment>>('put',`${BASE_URL}/v1/comment/reject/${id}`);
};

// 删除评论
export const deleteComment = (id: number) => {
  return http.request<ApiResponse<null>>('delete',`${BASE_URL}/v1/comment/delete/${id}`);
};
