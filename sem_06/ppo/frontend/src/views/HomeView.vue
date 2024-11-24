<template>
  <BaseSection>
    <div class="scrolling-component" ref="scrollComponent" v-if="allPosts.length">
      <PostCard  v-for="post in allPosts" :key="post.id" v-bind:post="post"/>
    </div>
    <div v-else>
      <EmptyCard :text="'No new posts'"/>
    </div>
  </BaseSection>
</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';

import { mapActions, mapGetters } from 'vuex';

import PostCard from '@/components/PostCard.vue';
import BaseSection from '@/components/BaseSection.vue';
import EmptyCard from '@/components/EmptyCard.vue';

export default defineComponent({
  name: 'HomeView',
  components: { EmptyCard, BaseSection, PostCard },
  data() {
    return {
      loading: true,
      bottom: false,
    };
  },
  mounted() {
    this.clearPosts();
    window.onscroll = () => {
      this.bottom = document.documentElement.scrollTop + window.innerHeight
        === document.documentElement.offsetHeight;
      console.error('heelloo', this.bottom);
      if (this.bottom) {
        this.loadMorePosts();
      }
    };
    this.getFeedPosts({ page: this.getPagePosts, num: this.getNumPosts });
    this.loading = false;
  },
  methods: {
    ...mapActions(['getFeedPosts', 'clearPosts', 'incPagePosts']),
    loadMorePosts() {
      if (this.totalPosts > this.getNumPosts * (this.getPagePosts - 1)) {
        this.loading = true;
        this.incPagePosts();
        this.getFeedPosts({ page: this.getPagePosts, num: this.getNumPosts });
        this.loading = false;
      }
    },
  },
  computed: {
    ...mapGetters(['allPosts', 'totalPosts', 'getPagePosts', 'getNumPosts']),
  },
});
</script>
