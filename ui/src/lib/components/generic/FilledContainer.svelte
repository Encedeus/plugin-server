<script lang="ts">
    export let size: "sm" | "md" | "lg" | "xl" = "sm";
    export let value = "";
    export let placeholder = "";
    export let className = "";
    export let label = "";
    export let error = false;
</script>

<div class="bg-indigo-800 rounded-lg {className}">
    <slot/>
</div>
