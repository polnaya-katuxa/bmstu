<template>
  <div class="card mb-4">
    <div class="card-body">
      <CardHeader :user="user"/>
      <UserCardDescription :text="user.description"/>
    </div>
    <DoubleCheckModal @deleted="deleteUser(user.login)" :id="'check'+user.login"/>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';
import DoubleCheckModal from '@/components/DoubleCheckModal.vue';
import CardHeader from '@/components/CardHeader.vue';
import UserCardDescription from '@/components/UserCardDescription.vue';

export default defineComponent({
  name: 'UserCard',
  components: { UserCardDescription, CardHeader, DoubleCheckModal },
  props: ['user'],
  methods: {
    ...mapActions(['deleteUser']),
  },
});
</script>

<style scoped>

</style>
