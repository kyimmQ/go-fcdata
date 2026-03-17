# Project Roadmap

## Overview

This document outlines the development roadmap for go-fcdata, including current status, planned features, and future enhancements.

---

## Version History

### v1.0.0 (Current)
**Status**: Completed

- [x] REST API client implementation
- [x] Authentication (Login)
- [x] Market data endpoints
  - [x] Securities list and details
  - [x] Index list and components
  - [x] OHLC (daily and intraday)
  - [x] Daily index and stock price
- [x] SignalR client for real-time streaming
- [x] Basic data models
- [x] Example application

---

## Planned Features

### v1.1.0 - Enhanced Data Types
**Target**: Q2 2026

#### REST API Enhancements
- [ ] Add more security detail fields (market cap, industry, etc.)
- [ ] Add trading statistics endpoints
- [ ] Add market overview data

#### Data Model Improvements
- [ ] Create strongly-typed OHLC struct with float values
- [ ] Add date/time parsing utilities
- [ ] Add validation for request parameters

### v1.2.0 - SignalR Enhancements
**Target**: Q3 2026

- [ ] Automatic reconnection on disconnect
- [ ] Connection state management (connected, connecting, disconnected)
- [ ] Subscribe/unsubscribe methods
- [ ] Multiple channel support
- [ ] Heartbeat/ping-pong handling

### v1.3.0 - Developer Experience
**Target**: Q3 2026

- [ ] Context support for all API calls
- [ ] Request/response logging middleware
- [ ] Rate limiting support
- [ ] Retry logic with exponential backoff
- [ ] Comprehensive test coverage

### v1.4.0 - Advanced Features
**Target**: Q4 2026

- [ ] Batch request support
- [ ] Caching layer for frequently accessed data
- [ ] WebSocket compression support
- [ ] Metrics/observability hooks
- [ ] Connection pool for high-volume scenarios

---

## Feature Requests

### High Priority
| Feature | Description | Status |
|---------|-------------|--------|
| Context Support | Pass context.Context to all API calls for cancellation and timeout | Planned |
| Error Types | Custom error types for different failure modes | Planned |
| Retry Logic | Automatic retry on transient failures | Planned |

### Medium Priority
| Feature | Description | Status |
|---------|-------------|--------|
| Reconnection | Automatic SignalR reconnection | Planned |
| Typed OHLC | Numeric types instead of strings | Planned |
| Logging | Request/response logging middleware | Planned |

### Low Priority
| Feature | Description | Status |
|---------|-------------|--------|
| Caching | In-memory cache for responses | Planned |
| Rate Limiting | Built-in rate limiting | Planned |
| Metrics | Prometheus/exportable metrics | Planned |

---

## API Coverage

### Completed Endpoints

| Endpoint | Method | Status |
|----------|--------|--------|
| `/api/v2/Market/AccessToken` | POST | Done |
| `/api/v2/Market/Securities` | GET | Done |
| `/api/v2/Market/SecuritiesDetails` | GET | Done |
| `/api/v2/Market/IndexList` | GET | Done |
| `/api/v2/Market/IndexComponents` | GET | Done |
| `/api/v2/Market/DailyOhlc` | GET | Done |
| `/api/v2/Market/IntradayOhlc` | GET | Done |
| `/api/v2/Market/DailyIndex` | GET | Done |
| `/api/v2/Market/DailyStockPrice` | GET | Done |

### Potential Future Endpoints
(Dependent on API availability)

- Historical index data
- Trading statistics
- Market depth/order book
- Company fundamentals
- News and announcements

---

## Technical Debt

### Current Items
- No custom error types
- No context.Context support
- Limited test coverage
- No retry logic
- No connection state tracking for SignalR

### Cleanup Tasks
- Refactor debug helpers into separate package
- Add integration tests
- Document all public APIs with examples

---

## Contribution Guidelines

### Development Process
1. Fork the repository
2. Create a feature branch
3. Implement changes with tests
4. Update documentation
5. Submit pull request

### Code Quality Standards
- All public APIs must have godoc comments
- Minimum 80% test coverage for new code
- No lint errors
- Follow Go coding conventions

---

## Dependencies Roadmap

### Current Dependencies
| Package | Version | Purpose |
|---------|---------|---------|
| gorilla/websocket | v1.5.3 | WebSocket client |

### Potential Future Dependencies
- `github.com/google/uuid` - UUID generation
- `github.com/patrickmn/go-cache` - Caching
- `github.com/prometheus/client_golang` - Metrics

---

## Release Process

### Version Numbering
- Major: Breaking changes
- Minor: New features (backward compatible)
- Patch: Bug fixes

### Release Checklist
- [ ] Update version in documentation
- [ ] Run all tests
- [ ] Update CHANGELOG.md
- [ ] Tag release in git
- [ ] Build and test example

---

## Community Requests

Track feature requests from users:
- Issue tracker: GitHub Issues
- Discussions: GitHub Discussions

---

## Timeline (Projected)

| Version | Target | Key Features |
|---------|--------|--------------|
| 1.0.0 | Done | Core functionality |
| 1.1.0 | Q2 2026 | Data type improvements |
| 1.2.0 | Q3 2026 | SignalR enhancements |
| 1.3.0 | Q3 2026 | DX improvements |
| 1.4.0 | Q4 2026 | Advanced features |

*Timeline subject to change based on user feedback and priorities.*