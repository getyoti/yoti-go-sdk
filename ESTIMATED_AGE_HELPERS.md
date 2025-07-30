# Estimated Age Helper Methods for Digital ID

This document describes the new helper methods for requesting the `estimated_age` attribute with automatic fallback to the `date_of_birth` attribute from Digital ID app shares.

## Overview

The new helper methods allow Relying Parties to easily request estimated age information from users while providing automatic fallback to date of birth when estimated age is not available. This enhances the user experience by requesting the most specific age information available while still providing age verification capabilities.

## Features

### 1. Policy Builder Helper Methods

#### Digital Identity Policy Builder (`digitalidentity` package)

- `WithEstimatedAge(options ...interface{})` - Requests estimated_age with date_of_birth fallback
- `WithEstimatedAgeOver(age int, options ...interface{})` - Age over verification with fallback
- `WithEstimatedAgeUnder(age int, options ...interface{})` - Age under verification with fallback

#### Dynamic Policy Builder (`dynamic` package)

- `WithEstimatedAge(options ...interface{})` - Requests estimated_age with date_of_birth fallback
- `WithEstimatedAgeOver(age int, options ...interface{})` - Age over verification with fallback
- `WithEstimatedAgeUnder(age int, options ...interface{})` - Age under verification with fallback

### 2. User Profile Helper Methods

#### Digital Identity User Profile (`digitalidentity` package)

- `EstimatedAge()` - Returns the estimated_age attribute if present
- `EstimatedAgeWithFallback()` - Returns estimated_age or falls back to date_of_birth

#### Profile User Profile (`profile` package)

- `EstimatedAge()` - Returns the estimated_age attribute if present
- `EstimatedAgeWithFallback()` - Returns estimated_age or falls back to date_of_birth

## Usage Examples

### Basic Estimated Age Request

```go
// Digital Identity
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithEstimatedAge().
    Build()

// Dynamic
policy, err := (&dynamic.PolicyBuilder{}).
    WithEstimatedAge().
    Build()
```

### Age Verification with Fallback

```go
// Request age over 18 verification with fallback
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithEstimatedAgeOver(18).
    Build()

// Request age under 21 verification with fallback
policy, err := (&dynamic.PolicyBuilder{}).
    WithEstimatedAgeUnder(21).
    Build()
```

### With Source Constraints

```go
constraint, err := (&digitalidentity.SourceConstraintBuilder{}).
    WithPassport("").
    Build()

policy, err := (&digitalidentity.PolicyBuilder{}).
    WithEstimatedAgeOver(18, constraint).
    Build()
```

### Retrieving Estimated Age from User Profile

```go
// Get estimated age directly
estimatedAge := userProfile.EstimatedAge()
if estimatedAge != nil {
    fmt.Printf("Estimated age: %s", estimatedAge.Value())
}

// Get estimated age with fallback
result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()
if result != nil {
    if isEstimatedAge {
        // estimated_age was returned
        estimatedAge := result.(*attribute.StringAttribute)
        fmt.Printf("Estimated age: %s", estimatedAge.Value())
    } else {
        // date_of_birth was returned as fallback
        dateOfBirth := result.(*attribute.DateAttribute)
        fmt.Printf("Date of birth: %s", dateOfBirth.Value().Format("2006-01-02"))
    }
}
```

### Combined Usage Example

```go
// 1. Create policy requesting estimated age with fallback
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithFullName().
    WithEstimatedAgeOver(18).
    WithEmail().
    Build()

if err != nil {
    log.Fatal(err)
}

// 2. Use policy to create share session
sessionSpec, err := (&digitalidentity.ShareSessionBuilder{}).
    WithPolicy(policy).
    WithCallbackUrl("https://yourdomain.com/yoti/callback").
    Build()

// ... create session and handle callback ...

// 3. Retrieve and use estimated age information
result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()
if result != nil {
    if isEstimatedAge {
        estimatedAge := result.(*attribute.StringAttribute)
        log.Printf("User estimated age: %s", estimatedAge.Value())
        // Use estimated age for business logic
    } else {
        dateOfBirth := result.(*attribute.DateAttribute)
        log.Printf("User date of birth: %s", dateOfBirth.Value().Format("2006-01-02"))
        // Calculate age from date of birth for business logic
    }
}
```

## Generated JSON Policy

When using the helper methods, the SDK generates JSON policies with the following structure:

### Basic Estimated Age Request

```json
{
  "wanted": [
    {
      "name": "estimated_age",
      "alternative_names": ["date_of_birth"],
      "accept_self_asserted": false
    }
  ],
  "wanted_auth_types": [],
  "wanted_remember_me": false
}
```

### Age Verification with Fallback

```json
{
  "wanted": [
    {
      "name": "estimated_age",
      "alternative_names": ["date_of_birth"],
      "derivation": "age_over:18",
      "accept_self_asserted": false
    }
  ],
  "wanted_auth_types": [],
  "wanted_remember_me": false
}
```

### With Constraints

```json
{
  "wanted": [
    {
      "name": "estimated_age",
      "alternative_names": ["date_of_birth"],
      "derivation": "age_over:18",
      "constraints": [
        {
          "type": "SOURCE",
          "preferred_sources": {
            "anchors": [
              {
                "name": "PASSPORT",
                "sub_type": ""
              }
            ],
            "soft_preference": false
          }
        }
      ],
      "accept_self_asserted": false
    }
  ],
  "wanted_auth_types": [],
  "wanted_remember_me": false
}
```

## Implementation Details

### Alternative Names Support

The helper methods automatically add `date_of_birth` as an alternative name to the `estimated_age` attribute request. This enables the backend to provide date of birth when estimated age is not available.

### Attribute Structure

The `WantedAttribute` structure now includes:

```go
type WantedAttribute struct {
    name               string
    alternativeNames   []string  // New field for fallback support
    derivation         string
    constraints        []constraintInterface
    acceptSelfAsserted bool
    Optional           bool
}
```

### Fallback Logic

The fallback logic is implemented in the user profile methods:

1. **EstimatedAgeWithFallback()** first attempts to retrieve the `estimated_age` attribute
2. If `estimated_age` is not found, it falls back to `date_of_birth`
3. Returns both the attribute and a boolean indicating which attribute was found
4. Returns `nil` if neither attribute is available

## Age Derivation Rules

When using age verification methods (`WithEstimatedAgeOver`, `WithEstimatedAgeUnder`), the SDK applies the standard age derivation rules:

- `age_over:18` - Verifies the user is 18 years old or older
- `age_under:21` - Verifies the user is under 21 years old

The derivation is applied to whichever attribute is available (estimated_age or date_of_birth).

## Error Handling

### Policy Building Errors

All helper methods follow the existing error handling patterns and will propagate errors through the policy builder's error mechanism.

### Profile Retrieval Errors

The `EstimatedAgeWithFallback()` method handles date parsing errors gracefully - if date_of_birth is present but cannot be parsed, it returns `nil`.

## Backward Compatibility

This implementation is fully backward compatible:

- Existing code continues to work unchanged
- New methods are additive and don't modify existing behavior
- JSON structure includes new fields only when using the new helper methods
- The `alternative_names` field uses `omitempty` to avoid affecting existing requests

## Testing

Comprehensive test coverage includes:

### Unit Tests
- Policy builder methods with and without constraints
- User profile attribute retrieval and fallback logic
- JSON marshaling and structure validation
- Error handling and edge cases

### Integration Tests
- End-to-end policy creation and usage
- Combined attribute requests
- Constraint application with estimated age

### Example Tests
- Demonstrate real-world usage patterns
- Show expected JSON output
- Validate API behavior

## Migration Guide

### From Existing Age Verification

If you're currently using:

```go
// Old approach
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithDateOfBirth().
    Build()
```

You can enhance it with:

```go
// New approach with fallback
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithEstimatedAge().
    Build()
```

### From Age Derivation

If you're currently using:

```go
// Old approach
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithAgeOver(18).
    Build()
```

You can enhance it with:

```go
// New approach with estimated age preference
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithEstimatedAgeOver(18).
    Build()
```

## Best Practices

1. **Use EstimatedAgeWithFallback()**: Always use the fallback method to handle both scenarios
2. **Check Return Type**: Always check which attribute was returned using the boolean flag
3. **Handle Both Types**: Implement logic to handle both estimated age strings and date of birth dates
4. **Apply Constraints Appropriately**: Use source constraints when you need specific verification levels
5. **Combine with Other Attributes**: The estimated age helpers work seamlessly with other attribute requests

## Future Enhancements

This implementation provides a foundation for future enhancements:

- Additional age-related attributes
- More sophisticated fallback logic
- Extended constraint support for age verification
- Age calculation utilities for date of birth fallbacks
