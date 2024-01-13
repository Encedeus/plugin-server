<script lang="ts">
    import * as stream from "stream";

    export let
        label: "before" | "after" = "before",
        defaultState: boolean = false,
        value: boolean = false,
        onclick: (() => void) | undefined = undefined;


</script>
<div>
    {#if label === "before"}
        <label>
            <slot/>
        </label>
    {/if}

    <input type="checkbox" bind:value={value} on:click={onclick} bind:checked={defaultState}/>

    {#if label === "after"}
        <label>
            <slot/>
        </label>
    {/if}
</div>