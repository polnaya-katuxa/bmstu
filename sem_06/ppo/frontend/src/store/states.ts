import * as openapi from '@/openapi/api';

export interface UsersState {
  users: Array<openapi.User>,
  totalUsers: number,
  page: number,
  num: number,
}

export interface UserState {
  user: openapi.User,
}

export interface Notification {
  errors: Array<string>,
}

export interface ProfileState {
  profile: openapi.User,
  subscribed: boolean,
  self: boolean,
}

export interface PostsState {
  posts: Array<openapi.Post>,
  totalPosts: number,
  page: number,
  num: number,
}

export interface CommentsState {
  post: openapi.Post,
  comments: Array<openapi.Comment>,
  totalComments: number,
  page: number,
  num: number,
}

export interface State {
  comments: CommentsState;
  posts: PostsState;
  profile: ProfileState;
  user: UserState;
  users: UsersState;
  notification: Notification;
}
