<template>
  <BaseSection>
    <PostCard :post="getPost"/>

    <div class="scrolling-component" ref="scrollComponent" v-if="allComments.length">
      <CommentCard  v-for="comm in allComments" :key="comm.id" v-bind:comm="comm" :post="getPost"
                    :login="curLogin"/>
    </div>

    <div v-else>
      <EmptyCard :text="'No comments'"/>
    </div>
    <CommentModal :opt="'comment'"/>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import CommentModal from '@/components/CommentModal.vue';
import PostCard from '@/components/PostCard.vue';
import CommentCard from '@/components/CommentCard.vue';
import EmptyCard from '@/components/EmptyCard.vue';

export default defineComponent({
  name: 'CommentsView',
  components: {
    EmptyCard,
    CommentCard,
    PostCard,
    CommentModal,
    BaseSection,
  },
  data() {
    return {
      loading: true,
      id: '',
      bottom: false,
    };
  },
  mounted() {
    this.clearComments();
    window.onscroll = () => {
      this.bottom = document.documentElement.scrollTop + window.innerHeight
        === document.documentElement.offsetHeight;
      if (this.bottom) {
        this.loadMoreComments();
      }
    };
    this.id = this.$route.params.postID.toString();

    this.getComments({ id: this.id, page: this.getPageComm, num: this.getNumComm });
    this.loading = false;
  },
  // created() {
  // },
  methods: {
    ...mapActions(['getComments', 'comment', 'uncomment', 'clearComments', 'incPageComments']),
    loadMoreComments() {
      if (this.totalComments > this.getNumComm * (this.getPageComm - 1)) {
        this.loading = true;
        this.incPageComments();
        this.getComments({ id: this.getPost.id, page: this.getPageComm, num: this.getNumComm });
        this.loading = false;
      }
    },
  },
  computed: {
    ...mapGetters(['allComments', 'curLogin', 'getPost', 'totalComments', 'getPageComm', 'getNumComm']),
  },
});
</script>

<style>
.card  {
  border-color: #d7c1f1 !important;
}
</style>
