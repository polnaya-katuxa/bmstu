<template>
  <div>
    <button v-if="location === `${basePath}users`" class="btn btn-sm btn-reaction" type="button"
            id="deleteButton" data-bs-toggle="modal"
            :data-bs-target="'#modal-' + 'check'+user.login"
            data-mdb-ripple-color="dark" style="z-index: 1;">
      Delete
    </button>

    <button class="btn btn-sm btn-reaction" type="button" id="deleteCommButton"
            v-else-if="location === `${basePath}comments/${post.id}` &&
              (curLogin === user.login || post.author.login === curLogin)"
            data-bs-toggle="modal" :data-bs-target="'#modal-' + 'check'+comm.id"
            data-mdb-ripple-color="dark" style="z-index: 1;">
      Delete
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters } from 'vuex';

export default defineComponent({
  name: 'CardHeaderButton',
  props: ['post', 'user', 'comm'],
  computed: {
    ...mapGetters(['curLogin']),
    isOwnProfile() {
      return (window.location.pathname === `${process.env.BASE_URL}profile/${this.curLogin}`);
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
