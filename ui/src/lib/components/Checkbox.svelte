<script lang="ts">

    export let
        label: "before" | "after" = "before",
        defaultState: boolean = false,
        value: boolean = false,
        onclick: (() => void) | undefined = undefined,
        className: string;


</script>
<div class={className}>
    {#if label === "before"}
        <label class="text-white font-bold text-[1rem] ml-4">
            <slot/>
        </label>
    {/if}

    <input type="checkbox" bind:value={value} on:click={onclick} bind:checked={defaultState}/>

    {#if label === "after"}
        <label class="text-white font-bold text-[1rem] ml-4">
            <slot/>
        </label>
    {/if}
</div>