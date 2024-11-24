<template>
  <BaseSection>
    <div class="scrolling-component" ref="scrollComponent" v-if="users.length">
      <UserCard v-for="user in users" :key="user.id" v-bind:user="user"/>
    </div>
    <div v-else>
      <EmptyCard :text="'No users'"/>
<!--      <div class="card">-->
<!--        <div class="card-body">-->
<!--          No users-->
<!--        </div>-->
<!--      </div>-->
    </div>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import UserCard from '@/components/UserCard.vue';
import EmptyCard from '@/components/EmptyCard.vue';

export default defineComponent({
  name: 'UsersView',
  components: { EmptyCard, UserCard, BaseSection },
  data() {
    return {
      loading: true,
      scrollComponent: ref(null),
      bottom: false,
    };
  },
  mounted() {
    this.clearUsers();
    window.onscroll = () => {
      this.bottom = document.documentElement.scrollTop + window.innerHeight
        === document.documentElement.offsetHeight;
      if (this.bottom) {
        this.loadMoreUsers();
      }
    };
    this.userInfo().then(() => {
      if (!this.isAdmin) {
        window.location.replace(process.env.BASE_URL || '');
      }

      ref(this.viewUsers({ page: this.getPageUsers, num: this.getNumUsers }));
      this.loading = false;
    });
  },
  methods: {
    ...mapActions(['viewUsers', 'deleteUser', 'userInfo', 'clearUsers', 'incPageUsers']),
    loadMoreUsers() {
      if (this.totalUsers > this.getNumUsers * (this.getPageUsers - 1)) {
        this.loading = true;
        this.incPageUsers();
        ref(this.viewUsers({ page: this.getPageUsers, num: this.getNumUsers }));
        this.loading = false;
      }
    },
  },
  computed: {
    ...mapGetters(['isAdmin', 'users', 'totalUsers', 'getPageUsers', 'getNumUsers']),
    location() {
      return window.location.pathname;
    },
  },
});
</script>

<style>
.card  {
  border-color: #d7c1f1 !important;
}
</style>
