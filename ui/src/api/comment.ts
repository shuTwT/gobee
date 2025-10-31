import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

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

// 评论列表响应类型
export interface CommentListResponse {
  data: Comment[];
  total: number;
  page: number;
  pageSize: number;
}

// 获取评论列表
export const getCommentList = (params: {
  page?: number;
  pageSize?: number;
  status?: CommentStatus;
}) => {
  return http.request<ApiResponse<CommentListResponse>>('get',`${BASE_URL}/v1/comments`, { params });
};

// 获取评论详情
export const getCommentDetail = (id: number) => {
  return http.request<ApiResponse<Comment>>('get',`${BASE_URL}/v1/comments/${id}`);
};

// 审核评论
export const approveComment = (id: number) => {
  return http.request<ApiResponse<Comment>>('put',`${BASE_URL}/v1/comments/${id}/approve`);
};

// 拒绝评论
export const rejectComment = (id: number) => {
  return http.request<ApiResponse<Comment>>('put',`${BASE_URL}/v1/comments/${id}/reject`);
};

// 删除评论
export const deleteComment = (id: number) => {
  return http.request<ApiResponse<null>>('delete',`${BASE_URL}/v1/comments/${id}`);
};
