<template>
  <button v-for="r in reactions" :key="r.typeID"
          :class="'btn btn-sm ' + reactionClass(r.yours)"
          @click="changeReaction({ postID: post.id, typeID: r.typeID })">
    <img v-bind:src="r.icon" class="me-2 reaction-icon" alt="">{{ r.num }}
  </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'CardFooterReactionButton',
  props: ['post'],
  methods: {
    reactionClass(reacted: boolean): string {
      if (reacted) {
        return 'btn-my-reaction';
      }

      return 'btn-reaction';
    },
    ...mapActions(['changeReaction']),
    ...mapGetters(['postReactions']),
  },
  computed: {
    reactions() {
      return this.post.reactions ?? [];
    },
  },
});
</script>

<style scoped>

</style>
