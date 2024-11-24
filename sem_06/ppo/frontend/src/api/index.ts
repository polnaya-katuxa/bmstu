import {
  Configuration, CommentApi, PostApi, UserApi, ReactionApi, SubscriberApi,
} from '@/openapi/';

export default {
  commentAPI: new CommentApi(
    new Configuration({
      basePath: 'http://postby.space/api',
      baseOptions: {
        withCredentials: true,
      },
    }),
  ),
  postAPI: new PostApi(
    new Configuration({
      basePath: 'http://postby.space/api',
      baseOptions: {
        withCredentials: true,
      },
    }),
  ),
  userAPI: new UserApi(
    new Configuration({
      basePath: 'http://postby.space/api',
      baseOptions: {
        withCredentials: true,
      },
    }),
  ),
  reactionAPI: new ReactionApi(
    new Configuration({
      basePath: 'http://postby.space/api',
      baseOptions: {
        withCredentials: true,
      },
    }),
  ),
  subscriberAPI: new SubscriberApi(
    new Configuration({
      basePath: 'http://postby.space/api',
      baseOptions: {
        withCredentials: true,
      },
    }),
  ),
};
