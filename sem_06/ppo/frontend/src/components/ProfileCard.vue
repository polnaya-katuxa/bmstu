<template>
  <div class="card">

    <ProfileCardHeader/>

    <ProfileCardBody/>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import ProfileCardHeader from '@/components/ProfileCardHeader.vue';
import ProfileCardBody from '@/components/ProfileCardBody.vue';

export default defineComponent({
  name: 'ProfileCard',
  components: {
    ProfileCardBody,
    ProfileCardHeader,
  },
  data() {
    return {
      loading: true,
      login: '',
      bottom: false,
    };
  },
  mounted() {
    this.clearPosts();
    window.onscroll = () => {
      this.bottom = document.documentElement.scrollTop + window.innerHeight
        === document.documentElement.offsetHeight;
      if (this.bottom) {
        this.loadMorePosts();
      }
    };
    this.login = this.$route.params.login.toString();

    this.viewProfile(this.login);
    this.getProfilePosts({ login: this.login, page: this.getPagePosts, num: this.getNumPosts });
    this.loading = false;
  },
  methods: {
    ...mapActions(['getProfilePosts', 'viewProfile', 'incPagePosts', 'clearPosts']),
    loadMorePosts() {
      if (this.totalPosts > this.getNumPosts * (this.getPagePosts - 1)) {
        this.loading = true;
        this.incPagePosts();
        this.getProfilePosts({ login: this.login, page: this.getPagePosts, num: this.getNumPosts });
        this.loading = false;
      }
    },
  },
  computed: {
    ...mapGetters(['totalPosts', 'getPagePosts', 'getNumPosts']),
  },
});
</script>
