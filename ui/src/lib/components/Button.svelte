<script lang="ts">

    export let
        onclick: (() => void) | undefined = undefined;
</script>

<button on:click={onclick} on><slot/></button>