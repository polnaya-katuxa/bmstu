<template>
  <div class="d-flex mb-3">

    <CardAvatar :pic="user.picture"/>

    <CardMainInfo :post="post" :user="user"/>

    <ActionSelection v-if="isOwnProfile" :post="post"/>

    <CardHeaderButton :user="user" :post="post" :comm="comm"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import ActionSelection from '@/components/ActionsSelection.vue';
import CardAvatar from '@/components/CardAvatar.vue';
import CardMainInfo from '@/components/CardMainInfo.vue';
import CardHeaderButton from '@/components/CardHeaderButton.vue';

export default defineComponent({
  name: 'CardHeader',
  components: {
    CardHeaderButton,
    CardMainInfo,
    CardAvatar,
    ActionSelection,
  },
  props: ['post', 'user', 'comm'],
  methods: {
    ...mapActions(['changePermsPost', 'deletePost']),
  },
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
