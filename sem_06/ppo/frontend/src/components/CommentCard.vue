<template>
  <div class="card mb-4">
    <div class="card-body">
      <CardHeader :post="post" :user="comm.commentator" :comm="comm"/>

      <CardContent :text="comm.content"/>
    </div>
    <DoubleCheckModal @deleted="this.uncomment({postID: post.id, commID: comm.id})"
                      :id="'check'+comm.id"/>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import DoubleCheckModal from '@/components/DoubleCheckModal.vue';
import CardContent from '@/components/CardContent.vue';
import CardHeader from '@/components/CardHeader.vue';

export default defineComponent({
  name: 'CommentCard',
  components: { CardHeader, CardContent, DoubleCheckModal },
  props: ['comm', 'post'],
  methods: {
    ...mapActions(['uncomment']),
  },
  computed: {
    ...mapGetters(['curLogin']),
  },
});
</script>

<style scoped>

</style>
