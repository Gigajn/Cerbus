# syntax=docker/dockerfile:1
FROM node:24-alpine AS deps
WORKDIR /repo
RUN corepack enable && corepack prepare pnpm@10.26.2 --activate

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./
COPY apps/portal/package.json apps/portal/package.json
COPY packages ./packages

RUN pnpm install --frozen-lockfile --filter ./apps/portal...

FROM node:24-alpine AS build
WORKDIR /repo
RUN corepack enable && corepack prepare pnpm@10.26.2 --activate
COPY --from=deps /repo/node_modules /repo/node_modules
COPY --from=deps /repo/packages /repo/packages
COPY apps/portal ./apps/portal
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./

RUN pnpm -C apps/portal build

FROM node:24-alpine AS runner
WORKDIR /repo
ENV NODE_ENV=production
RUN addgroup -S app && adduser -S app -G app
USER app

COPY --from=build /repo/apps/portal /repo/apps/portal
COPY --from=deps /repo/node_modules /repo/node_modules
COPY --from=deps /repo/packages /repo/packages
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./

EXPOSE 3000
CMD ["pnpm", "-C", "apps/portal", "start"]
