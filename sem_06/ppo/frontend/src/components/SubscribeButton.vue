<template>
  <button type="button" v-if="self" disabled class="btn btn-reaction"
          data-mdb-ripple-color="dark"
          style="z-index: 1;" @click.prevent="subscribe">
    Subscribe
  </button>
  <button type="button" v-else-if="subscribed" class="btn btn-reaction"
          data-mdb-ripple-color="dark"
          style="z-index: 1;" @click.prevent="subscribe">
    Unsubscribe
  </button>
  <button type="button" v-else class="btn btn-reaction"
          data-mdb-ripple-color="dark"
          style="z-index: 1;" @click.prevent="subscribe">
    Subscribe
  </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'SubscribeButton',
  methods: {
    ...mapActions(['subscribe']),
  },
  computed: {
    ...mapGetters(['subscribed', 'self']),
  },
});
</script>
