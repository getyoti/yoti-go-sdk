# Age Estimation Helper Methods

This document describes the age estimation functionality in the Yoti Go SDK.

## EstimatedAgeOver Method

For age estimation in the App, we expose only one method:

```go
EstimatedAgeOver(age int, buffer int, options ...interface{})
```

This method creates a policy with derivation `age_over:<age>:<buffer>` and `date_of_birth` in alternative names.

## Age Derivation Format

Age verification uses the format `age_over:age:buffer`:
- `estimated_age` verification: Checks if user is over `age + buffer`
- `date_of_birth` fallback: Checks if user is exactly over `age`

For example, `age_over:18:5` means:
- Estimated age verification: Check if user is over 23 (18 + 5)
- Date of birth fallback: Check if user is exactly over 18

## Features

### 1. Policy Builder Methods

#### Digital Identity Policy Builder (`digitalidentity` package)

- `EstimatedAgeOver(age int, buffer int, options ...interface{})` - Age over verification with buffer support and fallback

#### Dynamic Policy Builder (`dynamic` package)

- `EstimatedAgeOver(age int, buffer int, options ...interface{})` - Age over verification with buffer support and fallback

### 2. User Profile Helper Methods

#### Digital Identity User Profile (`digitalidentity` package)

- `EstimatedAge()` - Returns the estimated_age attribute if present
- `EstimatedAgeWithFallback()` - Returns estimated_age or falls back to date_of_birth

#### Profile User Profile (`profile` package)

- `EstimatedAge()` - Returns the estimated_age attribute if present
- `EstimatedAgeWithFallback()` - Returns estimated_age or falls back to date_of_birth

## Usage Examples

### Basic Usage (Age 18, Buffer 5)

```go
// Digital Identity
policy, err := (&digitalidentity.PolicyBuilder{}).
    WithFullName().
    WithEmail().
    EstimatedAgeOver(18, 5). // Estimated age checks for 23, date_of_birth fallback checks for 18
    Build()

// Dynamic
policy, err := (&dynamic.PolicyBuilder{}).
    WithFullName().
    WithEmail().
    EstimatedAgeOver(18, 5).
    Build()
```

### With Source Constraints

```go
constraint, err := (&digitalidentity.SourceConstraintBuilder{}).
    WithPassport("").
    Build()

policy, err := (&digitalidentity.PolicyBuilder{}).
    WithFullName().
    WithEmail().
    EstimatedAgeOver(18, 5, &constraint). // Estimated age checks for 23, date_of_birth fallback checks for 18
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
    EstimatedAgeOver(18, 5).
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

When using the EstimatedAgeOver method, the SDK generates JSON policies with the following structure:

### Age Verification with Fallback

```json
{
  "wanted": [
    {
      "name": "estimated_age",
      "alternative_names": ["date_of_birth"],
      "derivation": "age_over:18:5",
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
      "derivation": "age_over:18:5",
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

When using the EstimatedAgeOver method, the SDK applies the age derivation rule:

- `age_over:18:5` - Verifies the user is over 23 in estimated_age, or exactly over 18 in date_of_birth

The derivation is applied to whichever attribute is available (estimated_age or date_of_birth).

## Error Handling

### Policy Building Errors

The EstimatedAgeOver method follows the existing error handling patterns and will propagate errors through the policy builder's error mechanism.

### Profile Retrieval Errors

The `EstimatedAgeWithFallback()` method handles date parsing errors gracefully - if date_of_birth is present but cannot be parsed, it returns `nil`.

## Standard Recommendation

For all examples and implementations, please use:
- Age: **18**
- Buffer: **5**

This ensures consistent behavior across all applications using the Yoti SDK.

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
