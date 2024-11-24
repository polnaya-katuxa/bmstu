<template>
  <div v-if="location === `${basePath}comments/` + post.id">
    <button class="btn btn-sm btn-reaction" type="button" data-bs-toggle="modal"
            data-bs-target="#staticBackdrop5" data-mdb-ripple-color="dark"
            style="z-index: 1;" aria-expanded="false">
      Comment
    </button>
  </div>

  <div v-else>
    <router-link :to="'/comments/' + post.id">
      <button class="btn btn-sm btn-reaction"> {{ commentsNum }} comments </button>
    </router-link>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CardFooterCommentButton',
  props: ['post'],
  computed: {
    commentsNum() {
      return this.post.commentsNum ?? 0;
    },
    location() {
      return window.location.pathname;
    },
    basePath() {
      return process.env.BASE_URL;
    },
  },
});
</script>

<style scoped>

</style>
