<template>
  <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">
    <router-link style="color: black; text-decoration: none"
                 v-bind:to="'/profile/' + user.login">{{ user.login }}
    </router-link>

    <div class="text-muted small" v-if="location !== `${basePath}users`"
    >{{ new Date(post.pubTime).toLocaleString() }}</div>
    <div class="text-muted small" v-else>{{ user.mail }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CardMainInfo',
  props: ['user', 'post'],
  computed: {
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
