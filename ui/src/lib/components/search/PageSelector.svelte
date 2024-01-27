<script lang="ts">
    import {page} from "$app/stores";

    export let pagesTotal: number,
        currentPage: number,
        getAnchor: (number) => string;

    let pageAnchors: Array<{ href: string, pageNumber: number }>,
        hrefFirst: string,
        hrefLast: string;

    function getMiddleSegment(page: number): Array<number> {
        let start = page - 2;
        let end = page + 2;

        if (start < 1) {
            start = 1;
        }
        if (end > pagesTotal) {
            end = pagesTotal;
        }

        return Array.from({length: (end - start)}, (_, i) => start);
    }

    function getAnchors(pageNumbers): Array<string> {
        let anchors: Array<string>;

        pageNumbers.forEach((page: number) => {
            anchors.push(getAnchor(page));
        });

        return anchors;
    }

    function loadData() {
        const pageNumbers = getMiddleSegment(currentPage);
        const anchors = getAnchors(pageNumbers);

        let newPageAnchors: Array<{ href: string, pageNumber: number }> = [];

        for (let i = 0; i < pageNumbers.length; i++) {
            newPageAnchors.push({
                pageNumber: pageNumbers[i],
                href: anchors[i]
            });
        }

        pageAnchors = newPageAnchors;

        hrefFirst = getAnchor(1);
        hrefLast = getAnchor(pagesTotal);
    }

    $: $page.url.pathname && loadData();
    loadData();
</script>



<a href={hrefFirst}> 1 </a>

{#if currentPage > 4}
    <p>...</p>
{/if}

{#each pageAnchors as anchor}
    <a href={anchor.href}> {anchor.pageNumber} </a>
{/each}

{#if pagesTotal - currentPage >= 4}
    <p>...</p>
{/if}

{#if pagesTotal !== 1}
    <a href={hrefLast}> {pagesTotal} </a>
{/if}