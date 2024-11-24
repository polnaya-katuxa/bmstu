<template>
  <div class="mb-5">
    <ProfileAccentText :text="'About'"/>
    <ProfileDescriptionArea :text="profile.description"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters } from 'vuex';
import ProfileAccentText from '@/components/ProfileAccentText.vue';
import ProfileDescriptionArea from '@/components/ProfileDescriptionArea.vue';

export default defineComponent({
  name: 'ProfileDescription',
  components: { ProfileDescriptionArea, ProfileAccentText },
  computed: {
    ...mapGetters(['profile']),
  },
});
</script>
