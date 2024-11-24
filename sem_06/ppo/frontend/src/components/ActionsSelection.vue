<template>
  <div class="dropdown">
    <ActionSelectionButton :text="'Actions'"/>
    <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
      <li>
        <a class="dropdown-item small" href="#" v-if="post.perms"
           @click.prevent="changePermsPost(post.id)">Make free</a>
        <a class="dropdown-item small" href="#" v-else
           @click.prevent="changePermsPost(post.id)">Make paid</a>
      </li>
      <li><a class="dropdown-item small" href="#" data-bs-toggle="modal"
             :data-bs-target="'#modal-' + 'check'+post.id" data-mdb-ripple-color="dark"
             style="z-index: 1;">Delete</a></li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';
import ActionSelectionButton from '@/components/ActionSelectionButton.vue';

export default defineComponent({
  name: 'ActionSelection',
  components: { ActionSelectionButton },
  props: ['post'],
  methods: {
    ...mapActions(['changePermsPost']),
  },
});
</script>
