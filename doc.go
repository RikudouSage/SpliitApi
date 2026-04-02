// Package spliitapi is the module root for an unofficial Go client for the
// Spliit API.
//
// The module is split into focused packages:
//   - package spliit contains the main typed client and batching primitives,
//   - package endpoint contains endpoint descriptors,
//   - package shape contains request and response DTOs,
//   - package model contains domain models returned by the API,
//   - package amount contains helpers for fixed-point monetary values.
//
// Most applications will start with package spliit.
package spliitapi
