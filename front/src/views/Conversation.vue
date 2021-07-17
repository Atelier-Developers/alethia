<template>
<div>

</div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
  name: "Conversation",
  computed:{
    ...mapGetters("ConversationModules", ["conversations", "messages"]),
  },
  methods: {
    ...mapActions("ConversationModules", [
        "getConversations",
        "getMessages",
        "updateConversation",
        "deleteConversation",
        "createConversation",
        "createMessage"
    ]),
  },
  mounted() {
    this.getConversations();
  }
}
</script>

<style scoped>

</style>