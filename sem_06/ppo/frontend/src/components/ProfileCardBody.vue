<template>
  <div class="card-body p-4 text-black">

    <ProfileDescription/>

    <div class="d-flex justify-content-between align-items-center mb-4">
      <ProfileAccentText :text="'Posts'"/>
<!--      <p class="lead fw-normal mb-0">Posts</p>-->
      <PublishButton v-if="self"/>
    </div>

    <ProfilePostsContainer>
      <ProfilePosts/>
    </ProfilePostsContainer>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters } from 'vuex';
import ProfilePostsContainer from '@/components/ProfilePostsContainer.vue';
import ProfilePosts from '@/components/ProfilePosts.vue';
import PublishButton from '@/components/PublishButton.vue';
import ProfileDescription from '@/components/ProfileDescription.vue';
import ProfileAccentText from '@/components/ProfileAccentText.vue';

export default defineComponent({
  name: 'ProfileCardBody',
  components: {
    ProfileAccentText,
    ProfileDescription,
    PublishButton,
    ProfilePosts,
    ProfilePostsContainer,
  },
  computed: {
    ...mapGetters(['self']),
  },
});
</script>
