<template>
  <BaseSection>
    <div v-for="user in users" :key="user.id" class="card mb-4">
      <div class="card-body">
        <div class="d-flex mb-3">
          <div class="flex-shrink-0">
            <img v-bind:src="user.picture"
                 style="min-width: 50px; min-height: 50px; max-width: 50px;
                         max-height: 50px"
                 class="align-self-start mr-3 rounded-circle"
                 alt=""/>
          </div>
          <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">
            <router-link style="color: black; text-decoration: none"
                         v-bind:to="'/profile/' + user.login">{{ user.login }}
            </router-link>
            <div class="text-muted small">{{ user.mail }}
            </div>
          </div>

          <div>
            <button class="btn btn-sm btn-reaction" type="button" id="deleteButton"
                    @click.prevent="deleteUser(user.login)">
              Delete
            </button>
          </div>
        </div>

        <p>{{ user.description }}</p>
      </div>
    </div>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'UsersView',
  components: { BaseSection },
  mounted() {
    this.userInfo().then(() => {
      if (!this.isAdmin) {
        window.location.replace('/');
      }

      this.viewUsers();
    });
  },
  methods: {
    ...mapActions(['viewUsers', 'deleteUser', 'userInfo']),
  },
  computed: {
    ...mapGetters(['isAdmin', 'users']),
    location() {
      return window.location.pathname;
    },
  },
});
</script>

<style>
.card  {
  border-color: #d7c1f1 !important;
}
</style>
