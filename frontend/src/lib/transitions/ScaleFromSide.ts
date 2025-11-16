// src/lib/transitions/scaleFromSide.ts
import { cubicOut } from 'svelte/easing';

type Side = 'left' | 'right';
type Axis = 'x' | 'y' | 'both';

export interface ScaleFromSideParams {
  side?: Side;            // which horizontal edge anchors the scale
  start?: number;         // starting scale (0..1). e.g., 0.1 -> 10%
  axis?: Axis;            // scale axis (default 'x')
  duration?: number;      // ms
  easing?: (t: number) => number;
  fade?: boolean;         // also fade while scaling
}

export function scaleFromSide(
  node: Element,
  {
    side = 'left',
    start = 0,            // 0..1
    axis = 'x',
    duration = 200,
    easing = cubicOut,
    fade = false
  }: ScaleFromSideParams = {}
) {
  // clamp start to [0,1]
  start = Math.max(0, Math.min(1, start));

  const style = getComputedStyle(node);
  const base = style.transform === 'none' ? '' : style.transform;
  const origin = `${side} center`;
  const o0 = +style.opacity || 1;

  const scaleFn = (s: number) =>
    axis === 'y' ? `scaleY(${s})` :
    axis === 'both' ? `scale(${s})` :
    `scaleX(${s})`;

  return {
    duration,
    easing,
    css: (t: number) => {
      // t: 0→1 on intro, 1→0 on outro
      const s = start + (1 - start) * t;               // interpolate from start → 1
      const opacity = fade ? (o0 * t).toFixed(5) : o0.toString();

      return `
        transform-origin: ${origin};
        transform: ${base} ${scaleFn(s)};
        opacity: ${opacity};
        will-change: transform, opacity;
      `;
    }
  };
}
