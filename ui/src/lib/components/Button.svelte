<script lang="ts">
    import * as stream from "stream";

    export let
        onclick: (() => void) | undefined = undefined;
</script>

<button on:click={onclick} on><slot/></button>