<template>
  <div class="card-footer px-2">
    <div class="d-flex justify-content-between">
      <div>
        <div class="btn-group" role="group" aria-label="Reactions"
             v-if="reactions.length">

          <CardFooterReactionButton :post="post"/>

        </div> <p v-else>No reactions</p>
      </div>

      <CardFooterCommentButton :post="post"/>

    </div>
  </div>
</template>

<script lang="ts">
import CardFooterCommentButton from '@/components/CardFooterButton.vue';
import CardFooterReactionButton from '@/components/CardFooterReactionButton.vue';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CardFooter',
  components: { CardFooterReactionButton, CardFooterCommentButton },
  props: ['post'],
  computed: {
    reactions() {
      return this.post.reactions ?? [];
    },
  },
});
</script>
