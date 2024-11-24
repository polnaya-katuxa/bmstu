<template>
  <div class="rounded-top text-white d-flex flex-row" style="background-color: #fff;
            height:150px;">
    <div class="ms-4 mt-0 d-flex flex-column" style="width: 150px;">
      <ProfileCardAvatar :pic="profile.picture"/>

      <SubscribeButton/>

    </div>

    <ProfileCardHeaderMainInfo :textBigger="profile.login" :textSmaller="profile.mail"/>
  </div>

  <ProfileStatsBlock :stat1="totalPosts" :stat2="profile.balance"/>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters } from 'vuex';
import SubscribeButton from '@/components/SubscribeButton.vue';
import ProfileCardHeaderMainInfo from '@/components/ProfileCardHeaderMainInfo.vue';
import ProfileCardAvatar from '@/components/ProfileCardAvatar.vue';
import ProfileStatsBlock from '@/components/ProfileStatsBlock.vue';

export default defineComponent({
  name: 'ProfileCardHeader',
  components: {
    ProfileStatsBlock,
    ProfileCardAvatar,
    ProfileCardHeaderMainInfo,
    SubscribeButton,
  },
  computed: {
    ...mapGetters(['profile', 'totalPosts']),
  },
});
</script>
