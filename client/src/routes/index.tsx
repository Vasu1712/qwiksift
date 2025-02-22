import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
// import Infobox from "../components/starter/infobox/infobox";
import Image from '../media/qwiksift_name.svg?jsx';

export default component$(() => {
  return (
    <>
      <div class="flex flex-col gap-24 justify-center items-center p-8 py-24">
        <div class="flex justify-center text-2xl text-center ">
          <Image style={{ width: '300px'}} />
        </div>
        <div class="text-5xl text-center font-semibold bg-gradient-to-r from-grad1 via-grad2 via-grad3 via grad4 to-grad5 inline-block text-transparent bg-clip-text">
          Hello Vasu!
        </div>
        <div class="flex flex-col gap-1 justify-center text-center text-xl font-light text-midgray italic">
          <span>Your daily toolkit for fastest cheapest everything quick commerce</span>
          <span>Find your best deal with one simple search.</span>
        </div>
      </div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Qwiksift",
  meta: [
    {
      name: "Qwiksift Webapp",
      content: "Qwiksift is your one stop shop for all your daily needs. Compare amongst the best of quick commerce and brag the best deal within seconds.",
    },
  ],
};
