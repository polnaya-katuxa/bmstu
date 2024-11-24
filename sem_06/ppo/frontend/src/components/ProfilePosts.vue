<template>
  <div class="scrolling-component" ref="scrollComponent" v-if="allPosts.length">
    <PostCard v-for="post in allPosts" :key="post.id" v-bind:post="post"/>
  </div>

  <div v-else>
    <EmptyCard :text="'No posts'"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PostCard from '@/components/PostCard.vue';
import { mapGetters } from 'vuex';
import EmptyCard from '@/components/EmptyCard.vue';

export default defineComponent({
  name: 'ProfilePosts',
  components: {
    EmptyCard,
    PostCard,
  },
  computed: {
    ...mapGetters(['allPosts']),
  },
});
</script>
