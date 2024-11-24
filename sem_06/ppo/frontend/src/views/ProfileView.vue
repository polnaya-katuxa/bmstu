<template>
  <BaseSection>
    <ProfileCard/>
  </BaseSection>

  <PublishModal :opt="'post'"/>

</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions } from 'vuex';
import ProfileCard from '@/components/ProfileCard.vue';
import PublishModal from '@/components/PublishModal.vue';

export default defineComponent({
  name: 'ProfileView',
  components: { PublishModal, ProfileCard, BaseSection },
  mounted() {
    this.clearPosts();
  },
  created() {
    window.scrollTo(0, 0);
  },
  methods: {
    ...mapActions(['clearPosts']),
  },
});
</script>
