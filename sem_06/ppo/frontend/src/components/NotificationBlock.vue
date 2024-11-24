<template>
  <div class="notification-block">
    <div v-for="error in errors" :key="error"
         class="alert alert-danger alert-dismissible fade show" role="alert">
      <NotificationBody :error="error"/>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters } from 'vuex';
import NotificationBody from '@/components/NotificationBody.vue';

export default defineComponent({
  name: 'NotificationsBlock',
  components: { NotificationBody },
  computed: {
    ...mapGetters(['errors']),
  },
});
</script>
