<template>
  <div class="card-footer px-2">
    <div class="d-flex justify-content-between">
      <div>
        <div class="btn-group" role="group" aria-label="Reactions"
             v-if="reactions.length">

          <button v-for="r in reactions" :key="r.typeID"
                  :class="'btn btn-sm ' + reactionClass(r.yours)"
                  @click="changeReaction({ postID: post.id, typeID: r.typeID })">
            <img v-bind:src="r.icon" class="me-2 reaction-icon" alt="">{{ r.num }}
          </button>

        </div> <p v-else>No reactions</p>
      </div>

      <div v-if="location == '/comments/' + post.id">
        <button class="btn btn-sm btn-reaction" type="button" data-bs-toggle="modal"
                data-bs-target="#staticBackdrop5" data-mdb-ripple-color="dark"
                style="z-index: 1;" aria-expanded="false">
          Comment
        </button>
      </div>

      <div v-else>
        <router-link :to="'/comments/' + post.id">
          <button class="btn btn-sm btn-reaction"> {{ commentsNum }} comments </button>
        </router-link>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'ReactionsBlock',
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
      return this.post.reactions;
    },
    commentsNum() {
      return this.post.commentsNum;
    },
    location() {
      return window.location.pathname;
    },
  },
});
</script>

<style>
.reaction-icon {
  width: 22px !important;
  height: 22px !important;
}
.nav-comment {
  color: #8648c0;
}
.nav-comment:hover {
  font-weight: bolder;
}
</style>
