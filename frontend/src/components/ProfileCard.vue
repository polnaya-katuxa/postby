<template>
  <div class="card">
    <div class="rounded-top text-white d-flex flex-row" style="background-color: #fff;
            height:150px;">
      <div class="ms-4 mt-0 d-flex flex-column" style="width: 150px;">
        <img v-bind:src="profile.picture"
             alt="No image" class="img-fluid img-thumbnail mt-4 mb-2"
             style="width: 150px; min-height: 150px !important; z-index: 1">

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
      </div>
      <div class="ms-3" style="margin-top: 85px; color: #000;">
        <h5>{{ profile.login }}</h5>
        <p>{{ profile.mail }}</p>
      </div>
    </div>
    <div class="p-4 text-black" style="background-color: rgba(248, 242, 255, 0.99);">
      <div class="d-flex justify-content-end text-center py-1">
        <div>
          <p class="mb-1 h5">{{ allPosts.length }}</p>
          <p class="small text-muted mb-0">Posts</p>
        </div>
        <div class="px-3">
          <p class="mb-1 h5">{{ profile.balance }}</p>
          <p class="small text-muted mb-0">Coins</p>
        </div>
      </div>
    </div>
    <div class="card-body p-4 text-black">
      <div class="mb-5">
        <p class="lead fw-normal mb-1">About</p>
        <div class="p-4" style="background-color: rgba(248, 242, 255, 0.99);">
          <p class="font-italic mb-1">{{ profile.description }}</p>
        </div>
      </div>
      <div class="d-flex justify-content-between align-items-center mb-4">
        <p class="lead fw-normal mb-0">Posts</p>
        <button type="button" class="btn btn-reaction" data-bs-toggle="modal"
                data-bs-target="#staticBackdrop4" data-mdb-ripple-color="dark"
                style="z-index: 1;" v-if="self">
          Publish
        </button>
      </div>

      <div class="container w-746px">
        <br>
        <PostCard v-for="post in allPosts" :key="post.id" v-bind:post="post"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PostCard from '@/components/PostCard.vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'ProfileCard',
  components: {
    PostCard,
  },
  mounted() {
    const rl = this.$route.params.login.toString();

    this.viewProfile(rl);
    this.getProfilePosts(rl);
  },
  methods: {
    ...mapActions(['getProfilePosts', 'changePermsPost', 'deletePost',
      'viewProfile', 'subscribe']),
    reactionClass(reacted: boolean): string {
      if (reacted) {
        return 'btn-my-reaction';
      }

      return 'btn-reaction';
    },
  },
  computed: {
    ...mapGetters(['allPosts', 'profile', 'subscribed', 'self']),
  },
});
</script>
