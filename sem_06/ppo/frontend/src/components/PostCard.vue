<template>
  <div class="card mb-4">
    <div class="card-body">
      <CardHeader :post="post" :user="author"/>
      <CardContent :text="post.content"/>
    </div>

    <CardFooter :post="post"/>
    <DoubleCheckModal @deleted="this.deletePost(post.id)" :id="'check'+post.id"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CardFooter from '@/components/CardFooter.vue';
import { mapActions } from 'vuex';
import CardHeader from '@/components/CardHeader.vue';
import DoubleCheckModal from '@/components/DoubleCheckModal.vue';
import CardContent from '@/components/CardContent.vue';

export default defineComponent({
  name: 'PostCard',
  components: {
    CardContent,
    CardHeader,
    DoubleCheckModal,
    CardFooter,
  },
  props: ['post'],
  methods: {
    ...mapActions(['deletePost']),
  },
  computed: {
    author() {
      return this.post.author ?? {};
    },
  },
});
</script>
