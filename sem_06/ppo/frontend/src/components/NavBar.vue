<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <button class="navbar-toggler" type="button" data-toggle="collapse"
            data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false"
            aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav mx-auto">
        <li class="nav-item">
          <router-link class="nav-link px-2"  style="color: #fff;
                  text-decoration: underline" v-if="location==`${basePath}`"
                       to="/">Home</router-link>
          <router-link class="nav-link px-2"  style="color: #fff;
                  text-decoration: none" v-else to="/">Home</router-link>
        </li>
        <li class="nav-item">
          <a class="nav-link px-2" style="color: #fff; text-decoration: underline"
                       v-bind:href="`${basePath}profile/` + user.login"
                       v-if="location==`${basePath}profile/` + user.login"
          >Profile</a>
          <a class="nav-link px-2" style="color: #fff; text-decoration: none"
                       v-bind:href="`${basePath}profile/` + user.login" v-else>Profile</a>
        </li>
        <li class="nav-item">
          <router-link class="nav-link px-2" style="color: #fff; text-decoration: none"
                       @click.prevent="onLeave" to="/login">Leave
          </router-link>
        </li>
        <li class="nav-item" v-if="user.isAdmin">
          <router-link class="nav-link px-2" style="color: #fff; text-decoration: underline"
                       to="/users" v-if="location==`${basePath}users`">Users
          </router-link>
          <router-link class="nav-link px-2" style="color: #fff; text-decoration: none"
                       to="/users" v-else>Users
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script lang="ts">
import {
  User,
} from '@/openapi';
import Cookies from 'cookies-ts';
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'NavBar',
  mounted() {
    this.userInfo();
  },
  methods: {
    ...mapActions(['userInfo']),
    onLeave() {
      const cookies = new Cookies();
      cookies.remove('user-token');
      this.user = {} as User;
      window.location.replace(`${process.env.BASE_URL}login`);
    },
  },
  computed: {
    ...mapGetters(['user']),
    location() {
      return window.location.pathname;
    },
    basePath() {
      return process.env.BASE_URL;
    },
  },
});
</script>

<style>
.nav-link:hover {
  font-weight: bolder;
}
.nav-link {
  text-shadow: 1px 1px 3px rgba(0, 0, 0, 0.30);
}
</style>
