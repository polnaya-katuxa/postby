<template>
  <BaseSection>
    <div v-if="allPosts.length">
      <PostCard  v-for="post in allPosts" :key="post.id" v-bind:post="post"/>
    </div>
    <div v-else>
      <div class="card">
        <div class="card-body">
          No new posts
        </div>
      </div>
    </div>
  </BaseSection>
</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';

import { mapActions, mapGetters } from 'vuex';

import PostCard from '@/components/PostCard.vue';
import BaseSection from '@/components/BaseSection.vue';

export default defineComponent({
  name: 'HomeView',
  components: { BaseSection, PostCard },
  mounted() {
    this.getFeedPosts();
  },
  methods: {
    ...mapActions(['getFeedPosts']),
  },
  computed: {
    ...mapGetters(['allPosts']),
  },
});
</script>
