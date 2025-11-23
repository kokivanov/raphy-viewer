import { gsap } from "gsap";

type Params = {
  opacity?: number;
  duration?: number;
  ease?: string;
};

export function gsapFade(node: Element, params: Params = {}) {
  const {
    opacity = 0,
    duration = 0.4,
    ease = "power2.out"
  } = params;

  const tl = gsap.timeline({ paused: true });
  tl.fromTo(
    node,
    { opacity },
    { opacity: 1, duration, ease }
  );

  return {
    duration: duration * 1000,

    tick: (t: number) => {
      tl.progress(t);
    }
  };
}
